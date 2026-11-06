// structuredClone 是平台內建的深拷貝,不用自己遞迴複製。
export const clone = (config) => structuredClone(config);
