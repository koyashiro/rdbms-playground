import React from "react";
import Header from "~/components/Header.tsx";
import Playground from "~/components/Playground.tsx";

export default function Home() {
  return (
    <div className="App">
      <Header />
      <Playground />
    </div>
  );
}
