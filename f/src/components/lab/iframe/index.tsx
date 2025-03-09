'use client'

import React, { useState, useEffect } from 'react'
import { API_URL } from '@/config/config'

interface IframeProps {
  port: number
  path?: string
}

const Iframe: React.FC<IframeProps> = (props) => {
  const [count, setCount] = useState(0)
  const [host, setHost] = useState<string | undefined>(undefined)

  useEffect(() => {
    setHost(API_URL)
  }, [count])

  return (
    <>
      <button
        onClick={() => {
          setCount((count) => count + 1)
        }}
      >
        Reset
      </button>
      <iframe
        width={800}
        height={800}
        key={count}
        src={`http://${host}:${props.port}/${props.path}`}
      ></iframe>
    </>
  )
}

export default Iframe
