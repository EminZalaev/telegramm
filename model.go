package youtube

type RestResponse struct{
  Items []Item 'json:"items"'


}

type Item struct{
  id []ItemInfo 'json:"id"'

}

type ItemInfo struct{
  VideoId string 'json:"videoid"'

}
