import type { Meta, StoryObj } from '@storybook/react';

import { Tab } from '@/app/_components/organisms/Tab';

const itemChildren1 = (
  <div>
    ブラジルのコーヒーは、年間で200万～300万トンも生産されており、世界中のコーヒー豆生産量の約3分の1がブラジル産です。
    <br />
    コーヒー大国のブラジルの豆は日本でも広く普及している定番のコーヒーですね。
    <br />
    ブラジルのコーヒーは酸味やコクなどがバランスがとれた味わいが特徴的です。
  </div>
);

const itemChildren2 = (
  <div>
    グアテマラ産コーヒーはフルーティーな酸味と花のような香りが特長です。
    <br />
    深いコクがあり、やさしい甘さの後味が上品で飲みやすいでしょう。
    <br />
    産地によって違いがありますが、重めのボディでチョコレートやナッツのような甘みもあります。
    <br />
    重厚なボディのためブレンドのベースにも使われています。
  </div>
);

const itemChildren3 = (
  <div>
    国土の大半が山岳高原地帯で、収穫したコーヒーの運搬にはラバが使われることもあります。
    <br />
    そんなコロンビアで栽培されるコーヒー豆の全てがアラビカ種です。
    <br />
    コロンビア産のコーヒー豆は、甘い香りとしっかりした酸味とコク、重厚な風味で、バランスが良い特徴があります。
  </div>
);

const itemChildren4 = (
  <div>
    ルワンダコーヒーが特出するのは、レッドベリー、チョコレート、そして柑橘類を思わせる甘さと酸味。
    <br />
    国際的にもその品質は高く評価され、特にスペシャルティコーヒーとして魅力を放っています。
    <br />
    そして、彼らはただ美味いコーヒーを生産するだけでなく、
    <br />
    農家の生活向上、雇用の創出にも力を注いでいる社会貢献型のブランドでもあります。
  </div>
);

const meta: Meta<typeof Tab> = {
  component: Tab,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
  args: {},
};

export default meta;
type Story = StoryObj<typeof Tab>;

const textHeaderTabItems = [
  {
    id: '1',
    header: 'ブラジル',
    children: itemChildren1,
  },
  {
    id: '2',
    header: 'グアテマラ',
    children: itemChildren2,
  },
  {
    id: '3',
    header: 'コロンビア',
    children: itemChildren3,
  },
  {
    id: '4',
    header: 'ルワンダ',
    children: itemChildren4,
  },
];

export const TextHeaderTab: Story = {
  args: {
    defaultId: '1',
    items: textHeaderTabItems,
    width: 'w-[600px]',
    height: 'h-[300px]',
  },
};

const nodeHeaderTabItems = [
  {
    id: '1',
    header: <p>ブラジル</p>,
    children: itemChildren1,
  },
  {
    id: '2',
    header: <p>グアテマラ</p>,
    children: itemChildren2,
  },
  {
    id: '3',
    header: <p>コロンビア</p>,
    children: itemChildren3,
  },
  {
    id: '4',
    header: <p>ルワンダ</p>,
    children: itemChildren4,
  },
];

export const NodeHeaderTab: Story = {
  args: {
    defaultId: '1',
    items: nodeHeaderTabItems,
    width: 'w-[600px]',
    height: 'h-[300px]',
  },
};
