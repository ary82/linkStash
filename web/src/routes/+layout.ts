import type { LayoutLoad } from "./$types";

export const load: LayoutLoad = async ({ fetch }: any) => {
  try {
    const response = await fetch(`${import.meta.env.VITE_BACKEND_URL}/me`, {
      credentials: "include",
    });
    if (!response.ok) {
      throw new Error(`HTTP error: ${response.status}`);
    }
    const user = (await response.json()) as user;
    return { user };
  } catch (error) {
    console.error(error);
    throw new Error(`Unable to fetch`);
  }
};
