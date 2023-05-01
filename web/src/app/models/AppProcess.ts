import { Logs, State } from "./StreamProcess";

export interface AppProcess {
    name?:string
    docker_user?:string,
    docker_repository?:string,
    docker_ve