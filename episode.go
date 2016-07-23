package tvapi

import (
  "encoding/json"
  "fmt"
	"io"
	"os"
  "regexp"
  "strings"
)

type Episode struct {
  source string
  title string
  path string
  series *Series
  split []string
  season string
  episode string
}

func NewEpisode (source string, dbase *Dbase, config *Config) *Episode {
  var e Episode
  e.source = source
  series_index := e.digest()
  if series_index == 0 {
		return &e
	}

  tmp := strings.Join(e.split[0:series_index], "-")
  e.series = NewSeries(tmp, dbase, config)
  if e.series.api == 0 {
		return &e
	}

  e.Api()
  if e.title == "" {
		return &e
	}

	str := "%s/%s/%s-%s"
	e.path = fmt.Sprintf(str, config.video, e.series.key, e.series.key, e.season)
	if !e.check() {
		return &e
	}

  // TODO add cross mount checks for move
	dst := e.path + "/" + e.title
  mode := "copy"
  if config.move {
    mode = "move"
  }
	fmt.Println(mode, e.title)
	if config.move {
    err := os.Rename(e.source, dst)
    if err != nil {
      fmt.Println(err.Error())
    }
	} else {
    e.Copy()
	}

  return &e
}

func (e *Episode) Copy () bool {
	dst, err := os.Create(e.path + "/" + e.title)
  defer dst.Close()
  if err != nil {
		fmt.Println(err.Error())
    return false
  }

	src, err := os.Open(e.source)
  defer src.Close()
  if err != nil {
		fmt.Println(err.Error())
    return false
  }
  _, err = io.Copy(dst, src)
  if err != nil {
    fmt.Println(err.Error())
    return false
  }

  err = dst.Sync()
  if err != nil {
    fmt.Println(err.Error())
    return false
  }

	return true
}

func (e *Episode) Api () {
  url := "shows/%d/episodebynumber?season=%s&number=%s"
  url = fmt.Sprintf(url, e.series.api, e.season, e.episode)
  raw := util_curl(url)

	type Show struct {
    Name string
  }

  if len(raw) > 0 {
    var show Show
    err := json.Unmarshal(raw, &show)
    if err != nil {
			handle_err(err, 0)
			return
    }

		str := "%s.s%se%s.%s.%s"
		ext := e.split[len(e.split) - 1]
    e.title = util_format(show.Name, "episode")
		e.title = fmt.Sprintf(str, e.series.key, e.season, e.episode, e.title, ext)
  }
}

func (e *Episode) check () bool {
	dir, err := os.Open(e.path)
	defer dir.Close()
	if err != nil {
		if os.IsExist(err) {
			return false
		}

		err := os.MkdirAll(e.path, 0755)
		if err != nil {
			return false
		}
	}

	file, err := os.Open(e.path + "/" + e.title)
	defer file.Close()
	if err == nil {
		fmt.Println("xist", e.title, "skipping")
		return false
	}

	return true
}

func (e *Episode) digest () int {
  tmp := util_format(e.source, "episode")
  e.split = strings.Split(tmp, ".")

  for i, p := range e.split {
    if se, t := e.is_se(p); se {
      e.get_se(p, t)
      return i
      break
    }
  }

  return 0
}

func (e *Episode) get_se (se string, t int) {
  switch t {
  case 1:
    if len(se) == 3 {
      se = "0" + se
    }

    e.season = se[0:2]
    e.episode = se[2:4]
  break

  case 2:
    e.season = se[1:3]
    e.episode = se[4:6]
  break
  }
}

func (e *Episode) is_se (p string) (bool, int) {
  var r string
  var m bool

  r = "[0-9]{3,4}"
  if m, _ = regexp.MatchString(r, p); m {
    if len(p) == 3 {
      p = "0" + p
    }
    return true, 1
  }

  r = "s[0-9]{2}e[0-9]{2}"
  if m, _ = regexp.MatchString(r, p); m {
    return true, 2
  }

  return false, 0
}
