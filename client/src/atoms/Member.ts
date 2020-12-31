import { atom } from "recoil";

interface Member {
  name: string;
}

const defaultValue: Member[] = [];

export const MemberAtom = atom<Member[]>({
  key: "member-atom",
  default: defaultValue,
});
