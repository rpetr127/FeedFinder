const feeds = [];
const item = document.getElementsByClassName('item')
item[1].addEventListener("onclick", addItemToList)

function addItemToList() {
  console.log(item[0].textContent)
  const ind = feeds.length
  feeds[ind] = item[0].textContent
  return feeds
}
