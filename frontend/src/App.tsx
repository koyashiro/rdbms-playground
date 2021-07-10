import React from "react";
import "./App.css";
import Header from "./components/Header";
import Toolbar from "./components/Toolbar";
import Playground from "./components/Playground";

export default function App() {
  return (
    <div className="App">
      <Header />
      <Toolbar />
      <Playground />
    </div>
  );
}
