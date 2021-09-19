// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import React from 'react';

import { FormDescriptor } from './FormDescriptor';

export type FormProps = {
    form: FormDescriptor;
    children: React.ReactNode;
};

export const Form = (props: FormProps): JSX.Element => {
    const action = props.form.formSubmitAction;
    let URL = `${action.scheme}://${action.host}${action.path}`;
    if (action.rawQuery !== '') {
        URL += `?${action.rawQuery}`;
    }
    return (
        <form method={props.form.formSubmitMethod} action={URL}>
            { props.children }
        </form>
    );
};
