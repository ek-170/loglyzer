import '@testing-library/jest-dom';
import { render, screen } from '@/app/_test/test-utils';
import { Tab } from '../Tab';

const tabItems = [
  {
    id: '1',
    header: <p>Header1</p>,
    children: <div>Contents1</div>,
  },
  {
    id: '2',
    header: <p>Header2</p>,
    children: <div>Contents2</div>,
  },
  {
    id: '3',
    header: <p>Header3</p>,
    children: <div>Contents3</div>,
  },
  {
    id: '4',
    header: <p>Header4</p>,
    children: <div>Contents4</div>,
  },
];

describe('Tab', () => {
  it('tab rendering', async () => {
    const { user } = render(
      <>
        <Tab defaultId="1" items={tabItems} />
      </>,
    );
    const selectedHeader = screen.getByRole('button', { name: 'Header1' });
    expect(selectedHeader).toBeDisabled();
    expect(screen.getByText('Contents1')).toBeInTheDocument();

    await user.click(screen.getByRole('button', { name: 'Header2' }));
    expect(screen.getByText('Contents2')).toBeInTheDocument();
  });
});
