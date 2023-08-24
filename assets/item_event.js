const feeds = [];

(function () {
  const item = document.getElementsByClassName('item')
  item[1].Onlick(function () {
    const ind = feeds.length
    feeds[ind] = item
  })
  return feeds
})()
