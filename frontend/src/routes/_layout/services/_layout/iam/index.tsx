import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/_layout/services/_layout/iam/')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/_layout/services/iam/"!</div>
}
