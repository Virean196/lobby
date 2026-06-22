var ws;
var currentRoom = "";
function connect() {
  const username = document.getElementById("usernameInput").value;
  ws = new WebSocket(`ws://localhost:8080/ws?username=${username}`)
  ws.onopen = () => {
    console.log("Connected as " + username);
  }
  ws.onmessage = (event) => {
    const msg = JSON.parse(event.data);
    console.log(msg.type + ": " + msg.content);
  }

  ws.onclose = () => {
    console.log("Disconnected")
  }
}

function joinRoom() {
  const roomName = document.getElementById("roomInput").value;
  currentRoom = roomName;
  ws.send(JSON.stringify({ type: "join", room: roomName, content: "" }))
}

function sendMessage() {
  const message = document.getElementById("messageInput").value;
  ws.send(JSON.stringify({ type: "message", room: currentRoom, content: message }))
}
