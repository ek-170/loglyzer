const deleteAnalysis = async (url: string, { arg }: { arg: { username: string } }) => {
  const res = await fetch(url, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
  });
  return res.json()
}