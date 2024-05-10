import type { LoaderFunction, MetaFunction } from '@remix-run/node'
import XBlockDataTable from '~/components/home/blockDataTable'
import XMsgDataTable from '~/components/home/messageDataTable'
import Overview from '~/components/home/overview'
import { json } from '@remix-run/node'
import { gqlClient } from '~/entry.server'
import { useFetcher, useRevalidator, useSearchParams } from '@remix-run/react'
import { useInterval } from '~/hooks/useInterval'
import { xblockcount } from '~/components/queries/block'
import { xmsgs, xmsgrange } from '~/components/queries/messages'
import { XMsg } from '~/graphql/graphql'
import { mappedSourceChains } from '~/lib/sourceChains'
import { chainStats, supportedChains } from '~/components/queries/chains'

export const meta: MetaFunction = () => {
  return [
    { title: 'Omni Network Explorer' },
    { name: 'description', content: 'Omni Network Explorer' },
  ]
}

export type XmsgResponse = {
  supportedChainsList: any[]
  xmsgCount: string
  nextCursor: string
  prevCursor: string
  xmsgs: any[]
  pageInfo: {
    currentPage: string
    totalPages: string
    hasNextPage: boolean
    hasPrevPage: boolean
  },
  chainStats: any[]
}

type variablesPagination = {
  first?: number | null | undefined
  last?: number | null | undefined
  after?: string | null | undefined
  before?: string | null | undefined
  filters?: object[] | null
}

type variablesPagination = {
  first?: number | null | undefined
  last?: number | null | undefined
  after?: string | null | undefined
  before?: string | null | undefined
  filters?: object[] | null
}

export const loader: LoaderFunction = async ({ request }) => {
  // const res = await gqlClient.query(xblockcount, {})
  const url = new URL(request.url)
  const params: any = {}

  for (const [key, value] of url.searchParams) {
    params[key] = value
  }
  const variables: variablesPagination = {
    last: 10,
    filters: [],
  }

  if (params?.after) {
    variables['after'] = params?.after
    variables['first'] = 10
    variables['before'] = null
    variables['last'] = null
  }
  if (params?.before) {
    variables['before'] = params?.before
    variables['last'] = 10
    variables['after'] = null
    variables['first'] = null
  }
  if (params?.status) {
    if (params?.status === 'All') {
      variables['filters'] = []
    }

    if (params?.status === 'Success') {
      variables['filters']?.push({
        key: 'status',
        value: 'SUCCESS',
      })
    }
    if (params?.status === 'Pending') {
      variables['filters']?.push({
        key: 'status',
        value: 'PENDING',
      })
    }
    if (params?.status === 'Failed') {
      variables['filters']?.push({
        key: 'status',
        value: 'FAILED',
      })
    }
  }

  if (params?.address) {
    variables['filters']?.push({
      key: 'address',
      value: params?.address,
    })
  }
  if (params?.txHash) {
    variables['filters']?.push({
      key: 'txHash',
      value: params?.txHash,
    })
  }

  if (params.sourceChain) {
    variables['filters']?.push({
      key: 'srcChainID',
      value: params.sourceChain,
    })
  }

  if (params.destChain) {
    variables['filters']?.push({
      key: 'destChainID',
      value: params.destChain,
    })
  }
  if (params?.before) {
    variables['before'] = params?.before
    variables['first'] = null
    variables['after'] = null
    variables['last'] = 10
  }
  if (params?.status) {
    if (params?.status === 'All') {
      variables['filters'] = null
    }console.log(params?.status === 'Success');

    if (params?.status === 'Success') {
      variables['filters']?.push({
        key: 'status',
        value: 'SUCCESS',
      })
    }
    if (params?.status === 'Pending') {
      variables['filters']?.push({
        key: 'status',
        value: 'PENDING',
      })
    }
    if (params?.status === 'Failed') {
      variables['filters']?.push({
        key: 'status',
        value: 'FAILED',
      })
    }
  }

  console.log(variables);


  const [xmsgRes, supportedChainsRes, chainStatsRes] = await Promise.all([
    gqlClient.query(xmsgs, variables), // TODO when using the type generated by make gen, there's an issue where Type 'string' is not assignable to type 'never', type needs to be fixed when generated
    // gqlClient.query(xmsgrange, { from: '0x' + (0).toString(16), to: '0x' + (1000).toString(16) }),
    gqlClient.query(supportedChains, {}),
    gqlClient.query(chainStats, {}),
  ])

  const supportedChainsList = mappedSourceChains(supportedChainsRes.data?.supportedChains || [])
  const xmsgsList = xmsgRes.data?.xmsgs?.edges ?? []

  const pollData = async () => {
    return json({
      supportedChainsList: supportedChainsList,
      xmsgCount: xmsgRes.data?.xmsgs?.totalCount || 0,
      xmsgs: xmsgsList,
      pageInfo: xmsgRes.data?.xmsgs?.pageInfo,
      chainStats: chainStatsRes.data
    })
  }

  return await pollData()
}

export default function Index() {
  const revalidator = useRevalidator()

  return (
    <div className="px-4 sm:px-4 md:px-20  ">
      <div className="flex h-full w-full flex-col">
        <Overview />
        {/* <div className={'h-20'}></div> */}
        <div className="w-full">
          <XMsgDataTable />
        </div>
        <div className="grow"></div>
      </div>
    </div>
  )
}
