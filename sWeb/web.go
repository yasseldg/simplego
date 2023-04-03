package sWeb

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/yasseldg/simplego/sFile"
	"github.com/yasseldg/simplego/sLog"
	"github.com/yasseldg/simplego/sStr"

	"golang.org/x/net/html"
)

type DownloadsParams struct {
	Url      string
	Patterns []string
	Dir      string
	FileExt  []string
	Limit    int
}

func parse(url string) (*html.Node, error) {

	sLog.Info("Parse: %q", url)

	r, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("cannot get page")
	}

	b, err := html.Parse(r.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot parse page")
	}

	return b, err
}

func pageLinks(links []string, n *html.Node) []string {

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = pageLinks(links, c)
	}

	return links
}

// SearchLinks, use limit = -1 for nonstop
func SearchLinks(url, pathPrefix string, download bool, params *DownloadsParams, limit int) int {

	url_pp := url
	path_pp := params.Dir

	if len(pathPrefix) > 0 {
		url_pp = fmt.Sprintf("%s/%s", url_pp, pathPrefix)
		path_pp = fmt.Sprintf("%s/%s", path_pp, pathPrefix)
	}

	page, err := parse(url_pp)
	if err != nil {
		sLog.Error("SearchLinks: Error getting page: %q  ..  err: %s \n", url, err)
		return limit
	}

	links := pageLinks(nil, page)
	for _, link := range links {

		clearLink := strings.TrimRight(link, "/")

		file_url := fmt.Sprintf("%s/%s", url_pp, clearLink)

		if sStr.FindSuffix(link, params.FileExt...) {
			if sStr.FindPatterns(link, params.Patterns...) {
				if download {
					// log.Printf("Download: url: %s ", file_url)
					// log.Printf("Download: path: %s/%s ", path_pp, clearLink)

					err := DownloadFile(path_pp, clearLink, file_url)
					if err == nil {
						limit--
						sLog.Info("SearchLinks: Download success !!  file_url: %s ", file_url)
					} else {
						sLog.Error("SearchLinks: Error file_url: %s  ..  err: %s  ", file_url, err)
					}
				}
			}
		} else {
			// log.Printf("dir_url: %s \n", file_url)
			// log.Printf("dir_path: %s/%s \n\n", path_pp, clearLink)

			if len(pathPrefix) > 0 {
				d := false
				if sStr.FindPatterns(link, params.Patterns...) {
					d = download
				}
				limit = SearchLinks(url, fmt.Sprintf("%s/%s", pathPrefix, clearLink), d, params, limit)
			} else {
				limit = SearchLinks(url, clearLink, true, params, limit)
			}
		}

		if limit == 0 {
			break
		}
	}
	return limit
}

func DownloadFile(dir_path, file, url string) error {

	err := sFile.GetDir(dir_path)
	if err == nil {
		_, err = os.Stat(fmt.Sprintf("%s/%s", dir_path, file))
		if err != nil {
			if os.IsNotExist(err) {
				err = nil
				resp, err := http.Get(url)
				if err != nil {
					return err
				}
				defer resp.Body.Close()

				out, err := os.Create(fmt.Sprintf("%s/%s", dir_path, file))
				if err != nil {
					return err
				}
				defer out.Close()

				_, err = io.Copy(out, resp.Body)
				if err != nil {
					return err
				}
			} else {
				// Some other error. The file may or may not exist
				sLog.Error("DownloadFile: os.Stat( %q ): %s", fmt.Sprintf("%s/%s", dir_path, file), err)
			}
		} else {
			err = fmt.Errorf("DownloadFile: file exist ")
		}
	} else {
		sLog.Error("DownloadFile: Error GetDir( %s ): %s ", dir_path, err)
	}

	return err
}
