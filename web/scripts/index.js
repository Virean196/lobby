window.addEventListener("load", function() {
  var output = document.getElementById("output");
  var input = document.getElementById("input");
  var ws;

  function print(message) {
    var d = document.createElement("div");
    d.textContent = message;
    output.appendChild(d);
  }

  document.getElementById("open").onclick = function() {
    if (ws) return;
    ws = new WebSocket("ws://localhost:8080/echo");
    ws.onopen = function() {
      print("OPEN")
    }
    ws.onclose = () => {
      print("CLOSE");
      ws = null;
    };
    ws.onmessage = (evt) => print("RESPONSE: " + evt.data);
    ws.onerror = () => print("ERROR");
  };

  document.getElementById("send").onclick = function() {
    if (!ws) return;
    ws.send(input.value);
    print("SEND: " + input.value);
  };

  document.getElementById("close").onclick = function() {
    if (!ws) return;
    ws.close()
  };

});
