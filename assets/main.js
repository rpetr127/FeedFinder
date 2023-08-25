const feeds = require('/assets/item_event')

(function () {
  const refList = document.querySelector('p').querySelectorAll('a');
  const myInput = document.querySelector('input');

  for (let i = 0; i <= refList.length; i++) {
    refList[i].addEventListener('click', function (event) {
      event.preventDefault();
      myInput.value = event.target.text;
    })
  }
})();

var WebApp = window.Telegram.WebApp;

var MainButton = WebApp.MainButton;
var BackButton = WebApp.BackButton;

if !MainButton.isVisible {
  MainButton.show();
}
if !BackBtton.isVisible {
  BackButton.show();
}

MainButton.onClick(function () {
  const xhrURL = new URL('https://188.225.82.102:8443');
  for (let i = 0; i <= feeds.length; i++) {
    xhrURL.searchParams.set('feed', feeds[i]);
    const xhr = new XMLHttpRequest();
    xhr.open('GET', xhrURL);
    xhr.send();
    WebApp.sendData(feeds[i]);
  }
});
