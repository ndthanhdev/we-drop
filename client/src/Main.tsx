import React from "react";
import ReactDOM from "react-dom";
import {
  RecoilRoot,
  atom,
  selector,
  useRecoilState,
  useRecoilValue,
} from "recoil";
import Room from "./Room";

const MountNode = document.querySelector("#root");

function App() {
  return (
    <RecoilRoot>
      <div>
        hello world
        <Room />
      </div>
    </RecoilRoot>
  );
}

ReactDOM.render(<App />, MountNode);
