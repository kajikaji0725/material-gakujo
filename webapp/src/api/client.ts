import axios, { AxiosInstance } from "axios";
import { Seiseki, SeisekiJSON } from "./seiseki";

export class ApiClient {
  private baseURL: string;
  private client: AxiosInstance;

  constructor(baseURL: string) {
    this.baseURL = baseURL;
    const axiosClient = axios.create({
      baseURL: baseURL,
      withCredentials: true,
    });
    this.client = axiosClient;
  }

  async fetchSeisekis(): Promise<Seiseki[]> {
    const resp = await this.client.get<SeisekiJSON[]>("/seisekis");
    const seisekis = resp.data.map((s) => {
      return {
        ...s.Seiseki,
        id: s.ID,
        UserID: s.UserID,
        CreatedAt: s.CreatedAt,
        UpdatedAt: s.UpdatedAt,
      };
    });

    return seisekis;
  }

  async login(username: string, password: string): Promise<void> {
    const data = new URLSearchParams();
    data.set("gakujo_username", username);
    data.set("gakujo_password", password);
    const resp = await this.client.post<void>("/auth/login", data);
    return resp.data;
  }
}
