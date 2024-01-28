import type { Meta, StoryObj } from '@storybook/react';

import { LinkButton } from '@/app/_components/atoms/buttons/LinkButton';

const meta: Meta<typeof LinkButton> = {
  component: LinkButton,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof LinkButton>;

export const KibanaLinkButton: Story = {
  args: {
    color: 'elastic',
    padding: ['px-3', 'py-1.5'],
    width: 'w-8',
    href: '/',
    children: <p className="text-black">Kibana</p>,
  },
};
