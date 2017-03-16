package org.jsonstructure.jackson.validator;

import java.io.IOException;
import java.io.InputStream;
import java.util.HashMap;
import java.util.Map;

import javax.annotation.Nonnull;
import javax.annotation.Nullable;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import org.jsonstructure.jackson.validator.error.JSONStructureError;
import org.jsonstructure.jackson.validator.loanword.Result;

public class JSONStructureDefinition {

    @Nullable
    final String title;

    @Nullable
    final String description;

    @Nonnull
    final Map<String, String> imports;

    @Nonnull
    final Map<String, JsonNode> fragments;

    @Nonnull
    final Map<String, TypeDecl> types;

    @Nullable
    final TypeDecl main;

    @JsonCreator
    public JSONStructureDefinition(@JsonProperty("title") String title,
                                   @JsonProperty("description") String description,
                                   @JsonProperty("imports") Map<String, String> imports,
                                   @JsonProperty("fragments") Map<String, JsonNode> fragments,
                                   @JsonProperty("types") Map<String, TypeDecl> types,
                                   @JsonProperty("main") TypeDecl main) {
        this.title = title;
        this.description = description;
        this.imports = (imports == null) ? new HashMap<>() : imports;
        this.fragments = (fragments == null) ? new HashMap<>() : fragments;
        this.types = (types == null) ? new HashMap<>() : types;
        this.main = main;
    }

    @Nonnull
    public static Result<JSONStructureDefinition, JSONStructureError> create(@Nonnull InputStream inputStream)
            throws IOException {

        ObjectMapper objectMapper = Jackson.OBJECT_MAPPER;
        JsonNode tree = objectMapper.readTree(inputStream);
        JSONStructureError error = Compose.compose(tree);
        if (error != null) {
            return Result.err(error);
        }
        JSONStructureDefinition definition = objectMapper.treeToValue(tree, JSONStructureDefinition.class);
        return Result.ok(definition);
    }


}