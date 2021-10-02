const wordArray = CryptoJS.lib.WordArray.random(32);
const password = CryptoJS.enc.Base64.stringify(wordArray);

function currentPath() {
  return document.location.origin + document.location.pathname
}

function copyData(e) {
  var copyText = document.getElementById(e);
  copyText.select();
  copyText.setSelectionRange(0, 99999); /* For mobile devices */
  navigator.clipboard.writeText(copyText.value);
}

function writeURL(url) {
  var container = document.getElementById("container");
  var div = document.createElement("div");
  div.className = "secret-url"

  var h3 = document.createElement("h3");
  h3.textContent = "Copy URL to share secret"
  div.appendChild(h3);

  var p = document.createElement("p");
  p.textContent = "(link expires in 1 hour or upon viewing)"
  div.appendChild(p);

  var input = document.createElement("input");
  input.className = "secret-url";
  input.type = "text";
  input.id = "secret-url";
  input.value = url;
  div.appendChild(input);

  var button = document.createElement("button");
  button.onclick = function() {
    copyData("secret-url");
  };
  button.textContent = "Copy URL"
  div.appendChild(button);

  container.innerHTML = '';
  container.appendChild(div);
}

function encrypt(rawdata) {
  var ciphertext = CryptoJS.AES.encrypt(rawdata, password);
  return ciphertext.toString();
}

function decrypt() {
  var hash = window.location.hash;
  if (hash.length == 0) {
    return;
  }
  var password = hash.substr(1);
  var secretTextBox = document.getElementById("secret-text");
  var Normaltext = CryptoJS.AES.decrypt(secretTextBox.value, password);
  secretTextBox.textContent = Normaltext.toString(CryptoJS.enc.Utf8);
  secretTextBox.style.fontSize = "inherit";
}

function submitForm() {
  var secretTextBox = document.getElementById("secret-text");
  var data = {
    "secret": encrypt(secretTextBox.value)
  };

  var xhr = new XMLHttpRequest();
  var url = currentPath() + "s";
  xhr.open("POST", url, true);
  xhr.setRequestHeader("Content-Type", "application/json");
  xhr.onreadystatechange = function() {
    if (xhr.readyState === 4 && xhr.status === 200) {
      var json = JSON.parse(xhr.responseText);
      writeURL(currentPath() + "s/" + json.id + "#" + password);
    } else if (xhr.readyState === 4 && xhr.status === 413) {
      alert("Content-Length of posted data cannot exceed 100K");
      secretTextBox.value = "";
    }
  };
  xhr.send(JSON.stringify(data));
}

function copySecret() {
  copyData("secret-text");
}
