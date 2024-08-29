var port;

var deviceHygiene = { status: false, data: {} };

async function GetDeviceHygiene(pubKey) {
  function getChromeVersion() {
    var raw = navigator.userAgent.match(/Chrom(e|ium)\/([0-9]+)\./);
    return raw ? raw[2] : false;
  }

  let browserName = 'chrome';
  let browserVersion = '58';
  let build = '0';
  if (getChromeVersion() !== false) {
    browserName = 'chrome';
    browserVersion = getChromeVersion();
  } else {
    let brsr = await browser.runtime.getBrowserInfo();
    browserName = brsr.name;
    browserVersion = brsr.version;
    build = brsr.buildID;
  }

  let brsrDetails = {
    name: browserName,
    version: browserVersion,
    userAgent: navigator.userAgent,
    build: build,
  };

  let exts = await browser.management.getAll();
  let getHygieneReq = { pubKey: pubKey, deviceBrowser: brsrDetails, browserExtensions: exts };

  let message = { intent: 'getHygiene', data: getHygieneReq };

  chrome.runtime.sendNativeMessage('trasawrkstnagent', message, function(resp) {
    if (!browser.runtime.lastError) {
      // console.log('resp: ', JSON.parse(resp))
      return JSON.parse(resp.data);
    } else {
      onError(browser.runtime.lastError)
    }
  });
  
}

function onErr(e) {
  console.log('errr: ', e);
}

function onNativeMessage(message) {
  // console.log("[onNativeMessage] Received message: ", JSON.parse(message.response));
  deviceHygiene.status = true;
  deviceHygiene.data = message.response; //JSON.parse(message.response)
}

function onDisconnected() {
  //console.log("[onDisconnected] Disconnected: ", browser.runtime.lastError);
  deviceHygiene.status = false;
  deviceHygiene.data = {};
  port = null;
}
