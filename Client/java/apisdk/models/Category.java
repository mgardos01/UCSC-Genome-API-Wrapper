package apisdk.models;

import com.microsoft.kiota.serialization.ValuedEnum;
import java.util.Objects;

@jakarta.annotation.Generated("com.microsoft.kiota")
public enum Category implements ValuedEnum {
    HelpDocs("helpDocs"),
    PublicHubs("publicHubs"),
    TrackDb("trackDb");
    public final String value;
    Category(final String value) {
        this.value = value;
    }
    @jakarta.annotation.Nonnull
    public String getValue() { return this.value; }
    @jakarta.annotation.Nullable
    public static Category forValue(@jakarta.annotation.Nonnull final String searchValue) {
        Objects.requireNonNull(searchValue);
        switch(searchValue) {
            case "helpDocs": return HelpDocs;
            case "publicHubs": return PublicHubs;
            case "trackDb": return TrackDb;
            default: return null;
        }
    }
}
