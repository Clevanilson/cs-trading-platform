import type { HttpClient } from "./http-client";

export class HttpClientMock implements HttpClient {
  private _lastCalledEndpont = "";
  private _lastCalledBody: unknown = null;
  private _lastCalledMethod: "GET" | "POST" | "" = "";
  private _response: unknown = null;

  get lastCalledEndpont() {
    return this._lastCalledEndpont;
  }

  get lastCalledBody() {
    return this._lastCalledBody;
  }

  get lastCalledMethod() {
    return this._lastCalledMethod;
  }

  withResponse(response: unknown): this {
    this._response = response;
    return this;
  }

  async get<T = any>(url: string): Promise<T> {
    this._lastCalledEndpont = url;
    this._lastCalledMethod = "GET";
    this._lastCalledBody = null;
    return this._response as T;
  }

  async post<T = any>(url: string, body: unknown): Promise<T> {
    this._lastCalledEndpont = url;
    this._lastCalledMethod = "POST";
    this._lastCalledBody = body;
    return this._response as T;
  }
}
