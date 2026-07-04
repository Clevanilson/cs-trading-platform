import type { HttpClient } from "./http-client";

export class HttpClientFetchAdapter implements HttpClient {
  private readonly baseUrl: string = "http://localhost/api/";

  async get<T = any>(url: string): Promise<T> {
    const response = await fetch(`${this.baseUrl}${url}`);
    return response.json();
  }

  async post<T = any>(url: string, body: unknown): Promise<T> {
    const response = await fetch(`${this.baseUrl}${url}`, {
      method: "POST",
      body: JSON.stringify(body),
    });
    return response.json();
  }
}
