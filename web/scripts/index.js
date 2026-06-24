var ws;
var currentRoom = "";


function p(msg) {
  var output = document.getElementById("output")
  var node = document.createElement("div")
  node.textContent = msg
  output.appendChild(node)
}

function connect() {
  const username = document.getElementById("usernameInput").value;
  ws = new WebSocket(`ws://localhost:8080/ws?username=${username}`)
  ws.onopen = () => {
    console.log("Connected as " + username);
  }
  ws.onmessage = (event) => {
    const msg = JSON.parse(event.data);
    if (msg.type === "roster") {
      updateRoster(msg.content)
    } else if (msg.type === "message") {
      p(msg.sender + ": " + msg.content)
    }
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

function updateRoster(csv) {
  const users = csv.split(",")
  const roster = document.getElementById("roster")
  roster.innerHTML = ""
  users.forEach(name => {
    const div = document.createElement("div")
    div.textContent = name
    roster.appendChild(div)
  })
}

async function register() {
  const username = document.getElementById("usernameInputRegister").value
  const password = document.getElementById("passwordInput").value

  const response = await fetch("/register", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ username, password })
  })

  const data = await response.json()
  console.log(data)
}
