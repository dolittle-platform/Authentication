const params = new URLSearchParams(location.search);

$.get(
    '/.ory/kratos/public//self-service/errors?error='+params.get('error'),
    (data) => {
        const container = $('#container');
        for (const error of data.errors) {
            container.append(error.reason);
        }
    }
);
