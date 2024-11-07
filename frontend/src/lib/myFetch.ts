export interface BaseResponse<T> {
  ok: boolean;
  status: number;
  data: T;
  msg: string;
}

const baseURL = import.meta.env.VITE_API_URL;

export default async <T = any>(
  url: string,
  method: "POST" | "PUT" | "OPTIONS" | "GET" | "DELETE",
  reqData?: any
): Promise<BaseResponse<T>> => {
  const isFormData = reqData instanceof FormData;

  // Set headers
  const headers: RequestInit["headers"] = {
    "Content-Type": "application/json"
  };
  if (isFormData) {
    delete (headers as Record<string, string>)["Content-Type"];
  }

  // Set body and request URL
  let requsetUrl;
  if (["POST", "PUT"].includes(method)) {
    requsetUrl = `${baseURL}${url}`;
  } else {
    let params;
    if (params) {
      params = `?${new URLSearchParams(params)}`;
    }
    requsetUrl = `${baseURL}${url}${params || ""}`;
  }
  const body = reqData instanceof FormData ? reqData : JSON.stringify(reqData);

  // Set request options
  const req: RequestInit = {
    method,
    credentials: "include",
    body,
    headers
  };

  const res = await fetch(requsetUrl, req);
  const result = await res?.json();

  if (!res.ok) {
    console.log("Request failed: ", res?.statusText, "url: ", requsetUrl);
  }
  if (result.status === 400 && result.msg === "Invalid token") {
    console.log("Unauthorized");
    // clear token
    localStorage.removeItem("token");
    // redirect to login
    window.location.href = "/";
    return { ...result, ok: false };
  }
  return {
    ...result,
    ok: res.ok
  };
};
