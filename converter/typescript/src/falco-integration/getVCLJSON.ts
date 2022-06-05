import { exec as _exec } from "child_process";
import { promisify } from "util";
import { VCLAST } from "../ast/type";

const exec = promisify(_exec);

// TODO
const FALCO_COMMAND = `../../dist/falco-darwin-arm64 parse`;

export const getVCLJSON: (vclFiles: string) => Promise<VCLAST> = async function (vclFiles) {
  const jsonString = await exec(`${FALCO_COMMAND} ${vclFiles}`);
    return JSON.parse(jsonString.stdout);
};
