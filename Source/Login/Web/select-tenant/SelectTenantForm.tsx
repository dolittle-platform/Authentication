// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { useLocation } from 'react-router';

import { useResource } from 'rest-hooks';

import { Form } from '../forms/Form';
import { ChallengeNotFound } from './ChallengeNotFound';
import { Challenges } from './Challenges';
import { TenantList } from './TenantList';

import { Box } from '@mui/system';

const useLoginChallengeID = (): string | null => {
    const location = useLocation();
    const query = new URLSearchParams(location.search);
    return query.get('login_challenge');
};

export const SelectTenantForm = (): JSX.Element => {
    const challengeID = useLoginChallengeID();
    const challenge = useResource(Challenges, challengeID);

    if (challenge === null || challenge === undefined) throw new ChallengeNotFound(challengeID);

    return (
        <Form form={challenge.form}>
            <input type="hidden" name="login_challenge" value={challenge.id} />
            <Box sx={{ display: 'inline-block' }}>
                <TenantList tenants={challenge.user.tenants} />
            </Box>
        </Form>
    );
};
