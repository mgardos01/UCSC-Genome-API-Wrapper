<?php

namespace ApiSdk\Models;

use Microsoft\Kiota\Abstractions\Enum;

class Category extends Enum {
    public const HELP_DOCS = 'helpDocs';
    public const PUBLIC_HUBS = 'publicHubs';
    public const TRACK_DB = 'trackDb';
}
