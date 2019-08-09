window.onload = function () {
  handleBox(1);
  handleBox(2);
  handleBox(3);
  handleBox(4);
}

function handleBox(number) {
  // Get references to elements on the page.
  var form = document.getElementById('message-form' + number);
  var messageField = document.getElementById('message' + number);
  var messagesList = document.getElementById('messages' + number);
  var socketStatus = document.getElementById('status' + number);
  var closeBtn = document.getElementById('close' + number);


  // Create a new WebSocket.
  // var socket = new WebSocket('ws://echo.websocket.org');
  var socket = new WebSocket('ws://localhost:8000');


  // Handle any errors that occur.
  socket.onerror = function (error) {
    console.log('WebSocket Error: ' + error);
  };


  // Show a connected message when the WebSocket is opened.
  socket.onopen = function (event) {
    socketStatus.innerHTML = 'Connected to: ' + event.currentTarget.URL;
    socketStatus.className = 'open';
  };


  // Handle messages sent by the server.
  socket.onmessage = function (event) {
    var message = event.data;
    messagesList.innerHTML += '<li class="received"><span>Received:</span>' +
      message + '</li>';
  };


  // Show a disconnected message when the WebSocket is closed.
  socket.onclose = function (event) {
    socketStatus.innerHTML = 'Disconnected from WebSocket.';
    socketStatus.className = 'closed';
  };


  // Send a message when the form is submitted.
  form.onsubmit = function (e) {
    e.preventDefault();

    // Retrieve the message from the textarea.
    var message = messageField.value;

    // Send the message through the WebSocket.
    socket.send(message);

    // Add the message to the messages list.
    messagesList.innerHTML += '<li class="sent"><span>Sent:</span>' + message +
      '</li>';

    // Clear out the message field.
    messageField.value = '';

    return false;
  };


  // Close the WebSocket connection when the close button is clicked.
  closeBtn.onclick = function (e) {
    e.preventDefault();

    // Close the WebSocket.
    socket.close();

    return false;
  };

};