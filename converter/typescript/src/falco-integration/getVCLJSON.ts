import { exec as _exec } from "child_process";
import { promisify } from "util";

const exec = promisify(_exec);

// TODO
const FALCO_COMMAND = `../../dist/falco-darwin-arm64 parse`;

export const getVCLJSON: (vclFiles: string) => Promise<any> = async function (vclFiles) {
  const jsonString = await exec(`${FALCO_COMMAND} ${vclFiles}`);
    return JSON.parse(jsonString.stdout);
};
