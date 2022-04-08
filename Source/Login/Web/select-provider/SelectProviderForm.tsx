// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { useLocation } from 'react-router';

import { useResource } from 'rest-hooks';

import { Form } from '../forms/Form';
import { IdentityProviderList } from './IdentityProviderList';
import { Flows } from './Flows';
import { FlowNotFound } from './FlowNotFound';

const useFlowID = (): string | null => {
    const location = useLocation();
    const query = new URLSearchParams(location.search);
    return query.get('flow');
};

export const SelectProviderForm = (): JSX.Element => {
    const flowID = useFlowID();
    const flow = useResource(Flows, flowID);

    if (flow === null || flow === undefined) throw new FlowNotFound(flowID);

    return (
        <Form form={flow.form}>
            <IdentityProviderList providers={flow.providers} />
        </Form>
    );
};
