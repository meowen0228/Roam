export interface CustomRequestInit extends RequestInit {
  params?: any;
  onError?: (error: any) => void;
  onProgress?: (progress: number) => void;
}

export interface CustomBodyRequestInit extends RequestInit {
  body?: any;
  onError?: (error: any) => void;
  onProgress?: (progress: number) => void;
}

const baseURL = process.env.API_URL;

// 列出有可能是下載的檔案類型的Content-Type
const downloadContentType = [
  "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
  "application/vnd.ms-excel",
  "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
  "application/pdf",
  "application/zip",
  "application/octet-stream",
  "application/vnd.ms-powerpoint",
  "application/vnd.ms-outlook"
];

export const bodyFetch = async <T = any>(
  url: string,
  method: "POST" | "PUT",
  config?: CustomBodyRequestInit
): Promise<T> => {
  const accessToken = localStorage.getItem("accessToken");
  const reqData = config?.body;
  const headers: RequestInit["headers"] = {
    "Content-Type": "application/json",
    authorization: `Bearer ${accessToken}`,
    ...config?.headers
  };
  if (reqData instanceof FormData) {
    delete (headers as Record<string, string>)["Content-Type"];
  }
  const requsetUrl = `${baseURL}${url}`;
  const res = await fetch(requsetUrl, {
    method,
    credentials: "include",
    body: reqData instanceof FormData ? reqData : JSON.stringify(reqData),
    headers
  });
  const downloadRes = await downloadByResponse(res, config);
  if (downloadRes) {
    return downloadRes;
  }
  const result = await res?.json();
  if (!res.ok) {
    console.log("Request failed: ", res.statusText, "url: ", requsetUrl);
    config?.onError?.(result);
  }
  return result;
};

export const getFetch = async <T = any>(
  url: string,
  method: "GET" | "DELETE",
  config?: CustomRequestInit
): Promise<T> => {
  const accessToken = localStorage.getItem("accessToken");
  const headers = {
    authorization: `Bearer ${accessToken}`,
    ...config?.headers
  };
  let params;
  if (config?.params) {
    params = `?${new URLSearchParams(config.params)}`;
  }
  const requsetUrl = `${baseURL}${url}${params || ""}`;
  const res = await fetch(requsetUrl, {
    method,
    credentials: "include",
    headers
  });
  const downloadRes = await downloadByResponse(res, config);
  if (downloadRes) {
    return downloadRes;
  }
  const result = await res?.json();
  if (!res.ok) {
    console.log("Request failed: ", res.statusText, "url: ", requsetUrl);
    config?.onError?.(result);
  }
  return result;
};

// eslint-disable-next-line consistent-return
const downloadByResponse = async (res: Response, config?: CustomBodyRequestInit) => {
  //  判斷hearder是否是
  // @Header('Content-Type', 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet')
  // @Header('Content-Disposition', 'attachment; filename="xxxx.xlsx"')
  // 如果是就下載
  if (downloadContentType.includes(res.headers.get("Content-Type") ?? "")) {
    const filename = res.headers.get("Content-Disposition")?.split("filename=")[1].replace(/"/g, "");
    const contentLength = res.headers.get("Content-Length");
    const totalBytes = contentLength ? parseInt(contentLength, 10) : null;
    // 如果有Content-Length就顯示進度條

    if (totalBytes) {
      // 進度條
      console.log("totalBytes: ", totalBytes);
      const reader = res.body?.getReader();
      const stream = new ReadableStream({
        start(controller) {
          let receivedLength = 0;
          function push() {
            reader?.read().then(({ done, value }) => {
              if (done) {
                controller.close();
                return;
              }
              controller.enqueue(value); // 將讀取到的資料放入stream
              if (value && totalBytes) {
                receivedLength += value.length;
                const progress = receivedLength / totalBytes;
                config?.onProgress?.(progress);
              }
              setTimeout(push, 200);
            });
          }
          push();
        }
      });
      const newResponse = new Response(stream);
      const blob = await newResponse.blob();
      const aurl = window.URL.createObjectURL(blob);
      const a = document.createElement("a");
      a.href = aurl;
      a.download = filename || "download";
      a.click();
      window.URL.revokeObjectURL(aurl);
    }
    //   const chunks: Uint8Array[] = [];
    //   let receivedLength = 0;
    //   reader?.read().then(({ done, value }) => {
    //     if (value) {
    //       chunks.push(value);
    //       receivedLength += value.length;
    //       const progress = receivedLength / totalBytes;
    //       config?.onProgress?.(progress);
    //     }
    //   });
    //   const blob = new Blob(chunks);
    //   const aurl = window.URL.createObjectURL(blob);
    //   const a = document.createElement("a");
    //   a.href = aurl;
    //   a.download = filename || "download";
    //   a.click();
    //   window.URL.revokeObjectURL(aurl);
    //   return {
    //     status: res.status
    //   } as any;
    // }
    return {
      status: res.status
    } as any;
  }
};
