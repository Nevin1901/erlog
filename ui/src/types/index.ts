export interface ErLog {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt?: string;
  logType: string;
  message: string;
  title: string;
  extraData?: object;
}
