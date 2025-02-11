import { CenterLoader } from '@/components/common/loader'
import { useServer } from '@/hooks/useCompute'
import { createFileRoute } from '@tanstack/react-router'
import { useParams } from '@tanstack/react-router'

export const Route = createFileRoute(
  '/_layout/projects/$projectId/_layout/compute/server/$serverId/_layout/detail/',
)({
  component: RouteComponent,
})

function RouteComponent() {
  const { serverId } = useParams({ strict: false })
  const { isPending, isError, data, error } = useServer({
    id: serverId,
    refreshInterval: 10000,
  })

  console.log('DEBUG servers', isPending, isError, data, error)

  if (isPending) {
    return <CenterLoader />
  }

  if (isError) {
    return <div>Error: {error}</div>
  }

  if (data.error) {
    return <div>Error</div>
  }

  return <div>Detail</div>
}
