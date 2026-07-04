export interface HttpClient {
    get<T = any>(url: string): Promise<T>
    post<T = any>(url: string, body: unknown): Promise<T>
}