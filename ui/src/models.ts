export interface LogCount {
  id: number;
  title: string;
  message: string;
  num: number;
  logType: string;
}

export interface ILoaderData {
  logData: LogCount[];
  search: string;
}

export interface BaseModel {
  id: number;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string;
}

export interface ErLog extends BaseModel {
  logType: string;
  message: string;
  title: string;
  extraData?: any;
}

export interface IgnoredLog extends BaseModel {
  preview: string;
  hash: string;
}
