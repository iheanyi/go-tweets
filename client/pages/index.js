import React from 'react';
import MapGL from 'react-map-gl';

export default class extends React.Component {
  render() {
    return (
      <MapGL width={400} height={400} latitude={37.7577} longitude={-122.4376}
  zoom={8} onChangeViewport={(viewport) => {
    const {latitude, longitude, zoom} = viewport;
    // Optionally call `setState` and use the state to update the map.
  }}
  />);
  }
}
//export default () => <div>Hello World!</div>
