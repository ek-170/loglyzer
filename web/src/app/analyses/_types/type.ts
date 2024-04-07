export type Analysis = {
  id: string;
  // dataViewId: string,
  parseSources: ParseSource[];
};

export type ParseSource = {
  id: string; // ParseSource Info Doc ID
  name: string; // parse target file name
  index: string; // save target Index name
  order: number;
};
