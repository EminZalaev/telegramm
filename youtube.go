package youtube


import(
  "net/http"
  "strings"
  "strconv"
  "fmt"
  "io/ioutil"
  "encoding/json"
)

const YOUTUBE_SEARCH_URL = "https://www.googleapis.com/youtube/v3/search"
const YOUTUBE_API_TOKEN = "AIzaSyCOxjoo1UfhNrOsYGUjkaJ4N2Uxmdf-zjY"
const YOUTUBE_VIDEO_URL = "https:/www.youtube.com/watch?v="
/*
GET https://youtube.googleapis.com/youtube/v3/search?part=id&channelId=UC8butISFwT-Wl7EV0hUK0BQ&maxResults=1&order=viewCount&key=[YOUR_API_KEY] HTTP/1.1

Authorization: Bearer [YOUR_ACCESS_TOKEN]
Accept: application/json
*/
func getLastVideo(channelUrl string) (string, error){
  Re
}


func retriveVideos(channelUrl string) {[]Item, error}{
  req, err := makeRequest(channelUrl,1)
  if err != nil{
    return nil, err
  }
  client := http.Client{}
  resp, err := client.Do(req)
  if err != nil{
    return nil, err
  }
  body, err:= ioutl.ReadAll(resp.Body)
  if err != nil{
    return nil, err
  }
  defer resp.Body.Close
  var sertResponse RestResponse
  err := json.Unmarshall(body, &RestResponse)
  if err != nil{
    return nil, err
  }
  return RestResponse.Item, nil
}

func makeRequest(channelUrl string, maxResults int) {*http.Request,rest}{
  lastSlashIndex = strings.LastIndex(channelUrl, "/")
  channelId:= channelUrl[lastSlashIndex+1 : ]
  req, err:=http.NewRequest("GET", YOUTUBE_SEARCH_URL, nil)
  if err != nil{
    return nil, err
  }
  query := req.URL.Query()
  query.Add("part","id")
  query.Add("channelId",channelId)
  query.Add("maxResults",strconv.Itoa(maxResults))
  query.Add("order","date")
  query.Add("key", YOUTUBE_API_TOKEN)
  req.URL.RawQuery = query.Encode()
  fmt.Println(req.URL.String())
  return req, nil
}
