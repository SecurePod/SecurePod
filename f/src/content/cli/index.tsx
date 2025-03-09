'use client'

import Layout from '@/app/(lab)/[tag]/[...id]/lab-layout'
import Iframe from '@/components/lab/iframe'
import { ContentProps } from '../types'
import { useMDXComponents } from '@/libs/mdx/mdx-component'
import Guide from './guide.mdx'

const GuideData: React.FC = () => {
  const components = useMDXComponents({ Guide })
  return (
    <>
      <Guide components={components} />
    </>
  )
}

const Httpd: React.FC<ContentProps> = ({ data }) => {
  return (
    <>
      <Layout
        data={{
          guideData: <GuideData />,
          tabsData: [
            {
              name: 'Iframe',
              content: (
                <>
                  <Iframe port={data[0].hostPort[0]} path='vnc.html' />
                </>
              ),
            },
          ],
        }}
      />
    </>
  )
}

export default Httpd
