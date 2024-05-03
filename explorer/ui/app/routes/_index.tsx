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
import { supportedchains } from '~/components/queries/chains'
import { mappedSourceChains } from '~/lib/sourceChains'

export const meta: MetaFunction = () => {
  return [
    { title: 'Omni Network Explorer' },
    { name: 'description', content: 'Omni Network Explorer' },
  ]
}

export type XmsgResponse = {
  supportedChains: any[]
  xmsgCount: string
  nextCursor: string
  prevCursor: string
  xmsgs: any[]
}

export const loader: LoaderFunction = async ({ request }) => {
  // const res = await gqlClient.query(xblockcount, {})
  const url = new URL(request.url)
  const params: any = {}

  for (const [key, value] of url.searchParams) {
    params[key] = value
  }

  const variables = {
    limit: ('0x' + (10).toString(16)) as never,
  }
  if (params?.cursor) {
    variables['cursor'] = params?.cursor
  }

  const [xmsgRes, supportedChainsRes] = await Promise.all([
    gqlClient.query(xmsgs, variables), // TODO when using the type generated by make gen, there's an issue where Type 'string' is not assignable to type 'never', type needs to be fixed when generated
    // gqlClient.query(xmsgrange, { from: '0x' + (0).toString(16), to: '0x' + (1000).toString(16) }),
    gqlClient.query(supportedchains, {}),
  ])

  const supportedChains = mappedSourceChains(supportedChainsRes.data?.supportedchains || [])
  const xmsgsList = xmsgRes.data?.xmsgs?.Edges ?? []

  const pollData = async () => {
    return json({
      supportedChains,
      xmsgCount: xmsgRes.data?.xmsgs?.TotalCount || '0x0',
      nextCursor: xmsgRes.data?.xmsgs?.PageInfo?.NextCursor || null,
      prevCursor: xmsgRes.data?.xmsgs?.PageInfo?.PrevCursor || null,
      xmsgs: xmsgsList,
    })
  }

  return await pollData()
}

export default function Index() {
  const revalidator = useRevalidator()

  // poll server every 5 seconds
  // useInterval(() => {
  //   revalidator.revalidate()
  // }, 10000)

  return (
    <div className="px-8 md:px-20">
      <div className="flex h-full w-full flex-col">
        {/* <Overview /> */}
        <div className={'h-20'}></div>

        <div className="w-full">
          <XMsgDataTable />
        </div>
        <div className="grow"></div>
      </div>
    </div>
  )
}
