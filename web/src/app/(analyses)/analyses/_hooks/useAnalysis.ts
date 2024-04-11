import useSWR from 'swr'
import { fetcher }  from '@/app/_lib/fetch'
import useSWRMutation from 'swr/mutation';
import { Analysis } from '@/app/(analyses)/analyses/_types/type';

const createAnalysisFetcher = async (url: string, { arg }: { arg: { analysis: Omit<Analysis, "parseSources"> } }) => {
  const res = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
  });
  return res.json()
}

const deleteAnalysisFetcher = async (url: string) => {
  const res = await fetch(url, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
  });
  return res.json()
}

export function useAnalysis(id: string) {
  const { data, error, isLoading } = useSWR(`/api/analyses/${id}`, fetcher)

  const createAnalysis = useSWRMutation(`/api/analyses/${id}`, createAnalysisFetcher)

  // const updateAnalysis = async (analysis: Analysis) => {
  //   if (!data) {
  //     return false;
  //   }
  //   mutate();
  // };

  const deleteAnalysis = useSWRMutation(`/api/analyses/${id}`, deleteAnalysisFetcher)

  return {
    user: data,
    isLoading,
    isError: error,
    createAnalysis,
    // updateAnalysis,
    deleteAnalysis
  }
}