export const toLength = (text: string, len: number) => {
  return text.length > len ? text.substring(0, len - 3) + "..." : text;
};
