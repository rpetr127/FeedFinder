const feeds = [];

(function () {
  const item = document.getElementsByClassName('item')
  item[1].Onlick(function () {
    console.log(item[0].textContent)
    const ind = feeds.length
    feeds[ind] = item[0].textContent
  })
  return feeds
})()
