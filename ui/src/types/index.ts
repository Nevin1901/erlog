export interface ErLog {
  id: number;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string;
  logType: string;
  message: string;
  title: string;
  extraData?: object;
}
