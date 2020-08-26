import React from 'react';

export default class extends React.Component {
  componentDidMount() {
    console.log("In Component Did Mount Call!");
    if (window.WebSocket) {
      console.log("We got support for websockets!");
      let conn = new WebSocket("ws://localhost:1337/ws");

      conn.onopen = (evt) => {
        console.log("Opened Socket.");
        conn.send("This is an example message for stuff.");
      }; 

      conn.onerror = (evt) => {
        console.log("We got an error, folks.");
        console.log(evt);
        console.log(evt.data);
      };

      conn.onclose = (evt) => {
        console.log("Closing out connection!");
      };

      conn.onmessage = (evt) => {
        // Handle Messages here
        console.log(evt.data);
      };
    }
  }
  render() {
    return (
      <div>Hey Young World</div>
    );
  }
}
//export default () => <div>Hello World!</div>
