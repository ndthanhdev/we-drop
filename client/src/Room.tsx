import React from "react";
import { useRecoilValue } from "recoil";
import { MemberAtom } from "./atoms/Member";

let Room: React.FunctionComponent = () => {
  const memebers = useRecoilValue(MemberAtom).map((m) => <div>{m.name}</div>);

  return (
    <div>
      <div>room:</div>
      {memebers}
    </div>
  );
};

export default Room;
