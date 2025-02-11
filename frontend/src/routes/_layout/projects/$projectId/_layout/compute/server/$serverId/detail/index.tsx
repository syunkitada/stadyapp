import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute(
  '/_layout/projects/$projectId/_layout/compute/server/$serverId/detail/',
)({
  component: RouteComponent,
})

function RouteComponent() {
  return (
    <div>
      Hello
      "/_layout/projects/$projectId/_layout/compute/server/$serverId/detail/"!
    </div>
  )
}
