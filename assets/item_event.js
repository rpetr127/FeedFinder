var data = [];
const feeds_table = document.getElementById("feeds")
const items = feeds_table.children
for (var i = 0; i < feeds_table.children.length; i++) {
  const item = items[i].getElementsByTagName("td")
  item[1].addEventListener("click", function(event) {
    const text = item[0].textContent
    console.log(feedURL)
    const ind = data.length
    data[ind] = text 
    return data
  })
}
