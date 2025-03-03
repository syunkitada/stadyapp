// This file is auto-generated by @hey-api/openapi-ts

import type { Options as ClientOptions, TDataShape, Client } from '@hey-api/client-axios';
import type { GetWebUserData, GetWebUserResponse, GetWebUserError, GetPubkeysData, GetPubkeysResponse, GetPubkeysError, GetKeystoneVersionData, GetKeystoneVersionResponse, GetKeystoneVersionError, CreateKeystoneTokenData, CreateKeystoneTokenResponse, CreateKeystoneTokenError, CreateKeystoneFederationAuthTokenData, CreateKeystoneFederationAuthTokenResponse, CreateKeystoneFederationAuthTokenError, GetKeystoneDomainsData, GetKeystoneDomainsResponse, GetKeystoneDomainsError, CreateKeystoneDomainData, CreateKeystoneDomainResponse, CreateKeystoneDomainError, DeleteKeystoneDomainByIdData, DeleteKeystoneDomainByIdResponse, DeleteKeystoneDomainByIdError, GetKeystoneDomainByIdData, GetKeystoneDomainByIdResponse, GetKeystoneDomainByIdError, UpdateKeystoneDomainByIdData, UpdateKeystoneDomainByIdResponse, UpdateKeystoneDomainByIdError, UnassignKeystoneRoleFromUserDomainData, UnassignKeystoneRoleFromUserDomainResponse, UnassignKeystoneRoleFromUserDomainError, AssignKeystoneRoleToUserDomainData, AssignKeystoneRoleToUserDomainResponse, AssignKeystoneRoleToUserDomainError, GetKeystoneProjectsData, GetKeystoneProjectsResponse, GetKeystoneProjectsError, CreateKeystoneProjectData, CreateKeystoneProjectResponse, CreateKeystoneProjectError, DeleteKeystoneProjectByIdData, DeleteKeystoneProjectByIdResponse, DeleteKeystoneProjectByIdError, GetKeystoneProjectByIdData, GetKeystoneProjectByIdResponse, GetKeystoneProjectByIdError, UpdateKeystoneProjectByIdData, UpdateKeystoneProjectByIdResponse, UpdateKeystoneProjectByIdError, UnassignKeystoneRoleFromUserProjectData, UnassignKeystoneRoleFromUserProjectResponse, UnassignKeystoneRoleFromUserProjectError, AssignKeystoneRoleToUserProjectData, AssignKeystoneRoleToUserProjectResponse, AssignKeystoneRoleToUserProjectError, UnassignKeystoneRoleFromGroupProjectData, UnassignKeystoneRoleFromGroupProjectResponse, UnassignKeystoneRoleFromGroupProjectError, AssignKeystoneRoleToGroupProjectData, AssignKeystoneRoleToGroupProjectResponse, AssignKeystoneRoleToGroupProjectError, GetKeystoneUsersData, GetKeystoneUsersResponse, GetKeystoneUsersError, CreateKeystoneUserData, CreateKeystoneUserResponse, CreateKeystoneUserError, DeleteKeystoneUserByIdData, DeleteKeystoneUserByIdResponse, DeleteKeystoneUserByIdError, GetKeystoneUserByIdData, GetKeystoneUserByIdResponse, GetKeystoneUserByIdError, GetKeystoneUserProjectsByUserIdData, GetKeystoneUserProjectsByUserIdResponse, GetKeystoneUserProjectsByUserIdError, GetKeystoneGroupsData, GetKeystoneGroupsResponse, GetKeystoneGroupsError, GetKeystoneGroupByIdData, GetKeystoneGroupByIdResponse, GetKeystoneGroupByIdError, GetKeystoneRolesData, GetKeystoneRolesResponse, GetKeystoneRolesError, CreateKeystoneRoleData, CreateKeystoneRoleResponse, CreateKeystoneRoleError, DeleteKeystoneRoleByIdData, DeleteKeystoneRoleByIdResponse, DeleteKeystoneRoleByIdError, GetKeystoneRoleByIdData, GetKeystoneRoleByIdResponse, GetKeystoneRoleByIdError, UpdateKeystoneRoleByIdData, UpdateKeystoneRoleByIdResponse, UpdateKeystoneRoleByIdError, GetKeystoneRoleAssignmentsData, GetKeystoneRoleAssignmentsResponse, GetKeystoneRoleAssignmentsError, CreateKeystoneApplicationCredentialData, CreateKeystoneApplicationCredentialResponse, CreateKeystoneApplicationCredentialError } from './types.gen';
import { client as _heyApiClient } from './client.gen';

export type Options<TData extends TDataShape = TDataShape, ThrowOnError extends boolean = boolean> = ClientOptions<TData, ThrowOnError> & {
    /**
     * You can provide a client instance returned by `createClient()` instead of
     * individual options. This might be also useful if you want to implement a
     * custom client.
     */
    client?: Client;
};

/**
 * Returns user
 * Returns user
 *
 */
export const getWebUser = <ThrowOnError extends boolean = false>(options?: Options<GetWebUserData, ThrowOnError>) => {
    return (options?.client ?? _heyApiClient).get<GetWebUserResponse, GetWebUserError, ThrowOnError>({
        security: [
            {
                name: 'x-user-id',
                type: 'apiKey'
            }
        ],
        url: '/v1/web/user',
        ...options
    });
};

/**
 * Returns public keys
 * Returns public keys
 *
 */
export const getPubkeys = <ThrowOnError extends boolean = false>(options: Options<GetPubkeysData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).get<GetPubkeysResponse, GetPubkeysError, ThrowOnError>({
        url: '/v1/pubkeys',
        ...options
    });
};

/**
 * Get keystone version
 * Get keystone version
 */
export const getKeystoneVersion = <ThrowOnError extends boolean = false>(options?: Options<GetKeystoneVersionData, ThrowOnError>) => {
    return (options?.client ?? _heyApiClient).get<GetKeystoneVersionResponse, GetKeystoneVersionError, ThrowOnError>({
        url: '/keystone/v3',
        ...options
    });
};

/**
 * Create a new keystone token
 * Create a new keystone token
 */
export const createKeystoneToken = <ThrowOnError extends boolean = false>(options: Options<CreateKeystoneTokenData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).post<CreateKeystoneTokenResponse, CreateKeystoneTokenError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/auth/tokens',
        ...options,
        headers: {
            'Content-Type': 'application/json',
            ...options?.headers
        }
    });
};

/**
 * Create a new keystone token
 * Create a new keystone token
 */
export const createKeystoneFederationAuthToken = <ThrowOnError extends boolean = false>(options: Options<CreateKeystoneFederationAuthTokenData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).post<CreateKeystoneFederationAuthTokenResponse, CreateKeystoneFederationAuthTokenError, ThrowOnError>({
        security: [
            {
                name: 'x-user-id',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/OS-FEDERATION/identity_providers/{provider}/protocols/{protocol}/auth',
        ...options
    });
};

/**
 * Get keystone domains
 * Get keystone domains
 */
export const getKeystoneDomains = <ThrowOnError extends boolean = false>(options?: Options<GetKeystoneDomainsData, ThrowOnError>) => {
    return (options?.client ?? _heyApiClient).get<GetKeystoneDomainsResponse, GetKeystoneDomainsError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/domains',
        ...options
    });
};

/**
 * Create a new keystone domain
 * Create a new keystone domain
 */
export const createKeystoneDomain = <ThrowOnError extends boolean = false>(options: Options<CreateKeystoneDomainData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).post<CreateKeystoneDomainResponse, CreateKeystoneDomainError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/domains',
        ...options,
        headers: {
            'Content-Type': 'application/json',
            ...options?.headers
        }
    });
};

/**
 * Delete a domain by ID
 * Delete a domain by ID
 */
export const deleteKeystoneDomainById = <ThrowOnError extends boolean = false>(options: Options<DeleteKeystoneDomainByIdData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).delete<DeleteKeystoneDomainByIdResponse, DeleteKeystoneDomainByIdError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/domains/{id}',
        ...options
    });
};

/**
 * Get a domain by ID
 * Get a domain by ID
 */
export const getKeystoneDomainById = <ThrowOnError extends boolean = false>(options: Options<GetKeystoneDomainByIdData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).get<GetKeystoneDomainByIdResponse, GetKeystoneDomainByIdError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/domains/{id}',
        ...options
    });
};

/**
 * Update a domain by ID
 * Update a domain by ID
 */
export const updateKeystoneDomainById = <ThrowOnError extends boolean = false>(options: Options<UpdateKeystoneDomainByIdData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).patch<UpdateKeystoneDomainByIdResponse, UpdateKeystoneDomainByIdError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/domains/{id}',
        ...options,
        headers: {
            'Content-Type': 'application/json',
            ...options?.headers
        }
    });
};

/**
 * Unassign a role from a user on a domain
 * Unassign a role to a user on a domain
 */
export const unassignKeystoneRoleFromUserDomain = <ThrowOnError extends boolean = false>(options: Options<UnassignKeystoneRoleFromUserDomainData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).delete<UnassignKeystoneRoleFromUserDomainResponse, UnassignKeystoneRoleFromUserDomainError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/domains/{id}/users/{user_id}/roles/{role_id}',
        ...options
    });
};

/**
 * Assign a role to a user on a domain
 * Assign a role to a user on a domain
 */
export const assignKeystoneRoleToUserDomain = <ThrowOnError extends boolean = false>(options: Options<AssignKeystoneRoleToUserDomainData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).put<AssignKeystoneRoleToUserDomainResponse, AssignKeystoneRoleToUserDomainError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/domains/{id}/users/{user_id}/roles/{role_id}',
        ...options
    });
};

/**
 * Get keystone projects
 * Get keystone projects
 */
export const getKeystoneProjects = <ThrowOnError extends boolean = false>(options?: Options<GetKeystoneProjectsData, ThrowOnError>) => {
    return (options?.client ?? _heyApiClient).get<GetKeystoneProjectsResponse, GetKeystoneProjectsError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/projects',
        ...options
    });
};

/**
 * Create a new keystone project
 * Create a new keystone project
 */
export const createKeystoneProject = <ThrowOnError extends boolean = false>(options: Options<CreateKeystoneProjectData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).post<CreateKeystoneProjectResponse, CreateKeystoneProjectError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/projects',
        ...options,
        headers: {
            'Content-Type': 'application/json',
            ...options?.headers
        }
    });
};

/**
 * Delete a project by ID
 * Delete a project by ID
 */
export const deleteKeystoneProjectById = <ThrowOnError extends boolean = false>(options: Options<DeleteKeystoneProjectByIdData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).delete<DeleteKeystoneProjectByIdResponse, DeleteKeystoneProjectByIdError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/projects/{id}',
        ...options
    });
};

/**
 * Get a project by ID
 * Get a project by ID
 */
export const getKeystoneProjectById = <ThrowOnError extends boolean = false>(options: Options<GetKeystoneProjectByIdData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).get<GetKeystoneProjectByIdResponse, GetKeystoneProjectByIdError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/projects/{id}',
        ...options
    });
};

/**
 * Update a project by ID
 * Update a project by ID
 */
export const updateKeystoneProjectById = <ThrowOnError extends boolean = false>(options: Options<UpdateKeystoneProjectByIdData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).patch<UpdateKeystoneProjectByIdResponse, UpdateKeystoneProjectByIdError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/projects/{id}',
        ...options,
        headers: {
            'Content-Type': 'application/json',
            ...options?.headers
        }
    });
};

/**
 * Unassign a role from a user on a project
 * Unassign a role to a user on a project
 */
export const unassignKeystoneRoleFromUserProject = <ThrowOnError extends boolean = false>(options: Options<UnassignKeystoneRoleFromUserProjectData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).delete<UnassignKeystoneRoleFromUserProjectResponse, UnassignKeystoneRoleFromUserProjectError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/projects/{id}/users/{user_id}/roles/{role_id}',
        ...options
    });
};

/**
 * Assign a role to a user on a project
 * Assign a role to a user on a project
 */
export const assignKeystoneRoleToUserProject = <ThrowOnError extends boolean = false>(options: Options<AssignKeystoneRoleToUserProjectData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).put<AssignKeystoneRoleToUserProjectResponse, AssignKeystoneRoleToUserProjectError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/projects/{id}/users/{user_id}/roles/{role_id}',
        ...options
    });
};

/**
 * Unassign a role from a group on a project
 * Unassign a role to a group on a project
 */
export const unassignKeystoneRoleFromGroupProject = <ThrowOnError extends boolean = false>(options: Options<UnassignKeystoneRoleFromGroupProjectData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).delete<UnassignKeystoneRoleFromGroupProjectResponse, UnassignKeystoneRoleFromGroupProjectError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/projects/{id}/groups/{group_id}/roles/{role_id}',
        ...options
    });
};

/**
 * Assign a role to a group on a project
 * Assign a role to a group on a project
 */
export const assignKeystoneRoleToGroupProject = <ThrowOnError extends boolean = false>(options: Options<AssignKeystoneRoleToGroupProjectData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).put<AssignKeystoneRoleToGroupProjectResponse, AssignKeystoneRoleToGroupProjectError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/projects/{id}/groups/{group_id}/roles/{role_id}',
        ...options
    });
};

/**
 * Get keystone users
 * Create a new keystone token
 */
export const getKeystoneUsers = <ThrowOnError extends boolean = false>(options?: Options<GetKeystoneUsersData, ThrowOnError>) => {
    return (options?.client ?? _heyApiClient).get<GetKeystoneUsersResponse, GetKeystoneUsersError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/users',
        ...options
    });
};

/**
 * Create a new keystone user
 * Create a new keystone user
 */
export const createKeystoneUser = <ThrowOnError extends boolean = false>(options: Options<CreateKeystoneUserData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).post<CreateKeystoneUserResponse, CreateKeystoneUserError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/users',
        ...options,
        headers: {
            'Content-Type': 'application/json',
            ...options?.headers
        }
    });
};

/**
 * Delete a user by ID
 * Delete a user by ID
 */
export const deleteKeystoneUserById = <ThrowOnError extends boolean = false>(options: Options<DeleteKeystoneUserByIdData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).delete<DeleteKeystoneUserByIdResponse, DeleteKeystoneUserByIdError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/users/{id}',
        ...options
    });
};

/**
 * Get a user by ID
 * Get a user by ID
 */
export const getKeystoneUserById = <ThrowOnError extends boolean = false>(options: Options<GetKeystoneUserByIdData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).get<GetKeystoneUserByIdResponse, GetKeystoneUserByIdError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/users/{id}',
        ...options
    });
};

/**
 * Get user's projects
 * Get user's projects
 */
export const getKeystoneUserProjectsByUserId = <ThrowOnError extends boolean = false>(options: Options<GetKeystoneUserProjectsByUserIdData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).get<GetKeystoneUserProjectsByUserIdResponse, GetKeystoneUserProjectsByUserIdError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/users/{id}/projects',
        ...options
    });
};

/**
 * Get keystone groups
 * Create a new keystone token
 */
export const getKeystoneGroups = <ThrowOnError extends boolean = false>(options?: Options<GetKeystoneGroupsData, ThrowOnError>) => {
    return (options?.client ?? _heyApiClient).get<GetKeystoneGroupsResponse, GetKeystoneGroupsError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/groups',
        ...options
    });
};

/**
 * Get a group by ID
 * Get a group by ID
 */
export const getKeystoneGroupById = <ThrowOnError extends boolean = false>(options: Options<GetKeystoneGroupByIdData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).get<GetKeystoneGroupByIdResponse, GetKeystoneGroupByIdError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/groups/{id}',
        ...options
    });
};

/**
 * Get keystone roles
 * Get keystone roles
 */
export const getKeystoneRoles = <ThrowOnError extends boolean = false>(options?: Options<GetKeystoneRolesData, ThrowOnError>) => {
    return (options?.client ?? _heyApiClient).get<GetKeystoneRolesResponse, GetKeystoneRolesError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/roles',
        ...options
    });
};

/**
 * Create a new keystone role
 * Create a new keystone role
 */
export const createKeystoneRole = <ThrowOnError extends boolean = false>(options: Options<CreateKeystoneRoleData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).post<CreateKeystoneRoleResponse, CreateKeystoneRoleError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/roles',
        ...options,
        headers: {
            'Content-Type': 'application/json',
            ...options?.headers
        }
    });
};

/**
 * Delete a role by ID
 * Delete a role by ID
 */
export const deleteKeystoneRoleById = <ThrowOnError extends boolean = false>(options: Options<DeleteKeystoneRoleByIdData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).delete<DeleteKeystoneRoleByIdResponse, DeleteKeystoneRoleByIdError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/roles/{id}',
        ...options
    });
};

/**
 * Get a role by ID
 * Get a role by ID
 */
export const getKeystoneRoleById = <ThrowOnError extends boolean = false>(options: Options<GetKeystoneRoleByIdData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).get<GetKeystoneRoleByIdResponse, GetKeystoneRoleByIdError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/roles/{id}',
        ...options
    });
};

/**
 * Update a role by ID
 * Update a role by ID
 */
export const updateKeystoneRoleById = <ThrowOnError extends boolean = false>(options: Options<UpdateKeystoneRoleByIdData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).patch<UpdateKeystoneRoleByIdResponse, UpdateKeystoneRoleByIdError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/roles/{id}',
        ...options,
        headers: {
            'Content-Type': 'application/json',
            ...options?.headers
        }
    });
};

/**
 * Get keystone role assignments
 * Get keystone role assignments
 */
export const getKeystoneRoleAssignments = <ThrowOnError extends boolean = false>(options?: Options<GetKeystoneRoleAssignmentsData, ThrowOnError>) => {
    return (options?.client ?? _heyApiClient).get<GetKeystoneRoleAssignmentsResponse, GetKeystoneRoleAssignmentsError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/role_assignments',
        ...options
    });
};

/**
 * Create a new keystone application credential
 * Create a new keystone application credential
 */
export const createKeystoneApplicationCredential = <ThrowOnError extends boolean = false>(options: Options<CreateKeystoneApplicationCredentialData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).post<CreateKeystoneApplicationCredentialResponse, CreateKeystoneApplicationCredentialError, ThrowOnError>({
        security: [
            {
                name: 'x-auth-token',
                type: 'apiKey'
            }
        ],
        url: '/keystone/v3/users/{user_id}/application_credentials',
        ...options,
        headers: {
            'Content-Type': 'application/json',
            ...options?.headers
        }
    });
};