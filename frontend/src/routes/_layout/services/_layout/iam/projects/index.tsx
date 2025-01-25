import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/_layout/services/_layout/iam/projects/')(
  {
    component: RouteComponent,
  },
)

function RouteComponent() {
  return <div>Hello "/_layout/services/_layout/iam/projects/"!</div>
}
