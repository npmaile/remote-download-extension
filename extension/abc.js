const remDLSelection = "remote-download-selection";
const downloaderURL = "http://storage:5432/downloader";

async function triggerDownload(toDownload){
	console.log(JSON.stringify(toDownload))
	await fetch(downloaderURL,{
		method: 'POST',
		body: JSON.stringify(toDownload)
	})
}

function addMenuItem(){
	browser.menus.create({
		id:remDLSelection,
		title:"Download remotely"
	})
	browser.menus.refresh();
}

browser.menus.onClicked.addListener(async (info, tab) => {
  if (info.menuItemId === remDLSelection) {
	  await triggerDownload(info.linkUrl)
  }
});

browser.menus.onShown.addListener(function(info) {
	if (!info.linkUrl){
		return;
	}
	console.log(info.linkUrl);
	addMenuItem();
});

browser.menus.onHidden.addListener(function(){
	browser.menus.remove(remDLSelection);
});
